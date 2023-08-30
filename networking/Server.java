import java.io.IOError;
import java.io.IOException;
import java.net.ServerSocket;

public class Server {
    public static void main(String[] args) {
        try {
            ServerSocket server = new ServerSocket(80);
            server.accept();
            server.close();
        } catch (IOException e) {
            System.out.println(e);
        }    
    }
}