package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

const LOG_INFO = 0
const LOG_ERROR = -1

var memoryHashHistory []string

type ProcessInformation struct {
	PID         uint32
	ProcessName string
	ProcessPath string
	MemoryDump  []byte
}

// GetProcessesList return PID from running processes
func GetProcessesList() (procsIds []uint32, bytesReturned uint32, err error) {
	procsIds = make([]uint32, 2048)
	err = windows.EnumProcesses(procsIds, &bytesReturned)
	return procsIds, bytesReturned, err
}

// GetProcessModulesHandles list modules handles from a process handle
func GetProcessModulesHandles(procHandle windows.Handle) (processFilename string, modules []syscall.Handle, err error) {
	var processRawName []byte
	processRawName, err = GetProcessImageFileName(procHandle, 512)
	if err != nil {
		return "", nil, err
	}
	processRawName = bytes.Trim(processRawName, "\x00")
	processPath := strings.Split(string(processRawName), "\\")
	processFilename = processPath[len(processPath)-1]
	modules, err = EnumProcessModules(procHandle, 32)
	if err != nil {
		return "", nil, err
	}
	return processFilename, modules, nil
}

// DumpModuleMemory dump a process module memory and return it as a byte slice
func DumpModuleMemory(procHandle windows.Handle, modHandle syscall.Handle, verbose bool) []byte {
	moduleInfos, err := GetModuleInformation(procHandle, modHandle)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	memdump, err := ReadProcessMemory(procHandle, moduleInfos.BaseOfDll, uintptr(moduleInfos.SizeOfImage))
	if err != nil {
		log.Println("[ERROR]", err)
	}
	memdump = bytes.Trim(memdump, "\x00")
	return memdump
}

// GetProcessMemory return a process memory dump based on its handle
func GetProcessMemory(pid uint32, handle windows.Handle, verbose bool) (ProcessInformation, []byte, error) {
	procFilename, modules, err := GetProcessModulesHandles(handle)
	if err != nil {
		return ProcessInformation{}, nil, fmt.Errorf("Unable to get PID %d memory: %s", pid, err.Error())
	}
	for _, moduleHandle := range modules {
		if moduleHandle != 0 {
			moduleRawName, err := GetModuleFileNameEx(handle, moduleHandle, 512)
			if err != nil {
				return ProcessInformation{}, nil, err
			}
			moduleRawName = bytes.Trim(moduleRawName, "\x00")
			modulePath := strings.Split(string(moduleRawName), "\\")
			moduleFileName := modulePath[len(modulePath)-1]
			if procFilename == moduleFileName {
				return ProcessInformation{PID: pid, ProcessName: procFilename, ProcessPath: string(moduleRawName)}, DumpModuleMemory(handle, moduleHandle, verbose), nil
			}
		}
	}
	return ProcessInformation{}, nil, fmt.Errorf("Unable to get PID %d memory: no module corresponding to process name", pid)
}

// WriteProcessMemoryToFile try to write a byte slice to the specified directory
func WriteProcessMemoryToFile(path string, file string, data []byte) (err error) {
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0600); err != nil {
			return err
		}
	}
	if err := os.WriteFile(path+"/"+file, data, 0644); err != nil {
		return err
	}

	return nil
}

func logMessage(logType int, logMessage ...interface{}) {
	if logType == LOG_INFO {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(os.Stderr)
	}

	log.Println(logMessage...)
}

// GetProcessHandle return the process handle from the specified PID
func GetProcessHandle(pid uint32, desiredAccess uint32) (handle windows.Handle, err error) {
	handle, err = windows.OpenProcess(desiredAccess, false, pid)
	return handle, err
}

// StringInSlice check wether or not a string already is inside a specified slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ListProcess try to get all running processes and dump their memory, return a ProcessInformation slice
func ListProcess(verbose bool) (procsInfo []ProcessInformation) {
	runningPID := os.Getpid()

	procsIds, bytesReturned, err := GetProcessesList()
	if err != nil {
		log.Fatal(err)
	}
	for i := uint32(0); i < bytesReturned; i++ {
		if procsIds[i] != 0 && procsIds[i] != uint32(runningPID) {
			procHandle, err := GetProcessHandle(procsIds[i], windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ)
			if err != nil && verbose {
				logMessage(LOG_ERROR, "[ERROR]", "PID", procsIds[i], err)
			}

			if err == nil && procHandle > 0 {
				if proc, memdump, err := GetProcessMemory(procsIds[i], procHandle, verbose); err != nil && verbose {
					logMessage(LOG_ERROR, "[ERROR]", err)
				} else {
					if len(memdump) > 0 {
						// return process memory only if it has changed since the last process scan
						if !StringInSlice(fmt.Sprintf("%x", md5.Sum(memdump)), memoryHashHistory) {
							proc.MemoryDump = memdump
							procsInfo = append(procsInfo, proc)
							memoryHashHistory = append(memoryHashHistory, fmt.Sprintf("%x", md5.Sum(memdump)))
						}
					}
				}

			}
			windows.CloseHandle(procHandle)
		}
	}
	return procsInfo
}

func main() {
	fmt.Println(GetProcessesList())
}
