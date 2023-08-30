import javax.swing.*;

public class HiLo {
    public static void main(String[] args) {
        JFrame window = new JFrame("HiLo");
        JPanel panel = new JPanel();
        JButton button = new JButton("OK");
        panel.add(button);
        window.add(panel);
        window.setSize(300,100);
        button.addActionListener(e ->
           System.out.println("Mouse click!")
        ); 
        window.setVisible(true);
    }
}