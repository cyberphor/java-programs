# C Sharp
### How to Compile C Sharp Source Code into Executable Files
**Step 1.** Using PowerShell, create an alias for the C Sharp compiler (`csc.exe`). 
```pwsh
Set-Alias -Name "SendTo-CSharpCompiler" -Value "C:\windows\Microsoft.NET\Framework\v4.0.30319\csc.exe"
```

**Step 2.** Write your source code. 
```cs
// file-name: HelloWorld.cs

using System;

class HelloWorld {
    static void Main(string[] args) {
        Console.WriteLine("Hello world!");
    }
}
```

**Step 3.** Compile, execute, and profit! 
```pwsh
SendTo-CSharpCompiler HelloWorld.cs
```
```pwsh
.\HelloWorld.exe
```
```pwsh
Hello world!
```