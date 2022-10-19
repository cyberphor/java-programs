package java.searchThis;
import java.awt.*;
import javax.swing.*;

// class 
public class SearchThis{
    // class constructor
    SearchThis() {
        // declare a frame to house the buttons
        JFrame windowFrame = new JFrame();

        // add button1
        JButton button1 = new JButton("Date");
        windowFrame.add(button1);   

        // add button2
        JButton button2 = new JButton("Amount");
        windowFrame.add(button2);

        // add text box
        JTextField textField = new JTextField("Search");
        windowFrame.add(textField);

        // set layout, size, and visbility
        windowFrame.setLayout(new FlowLayout());
        windowFrame.setSize(300, 300);    
        windowFrame.setVisible(true);
    }

    // main
    public static void main(String args[]){
        // calls for a new instance of the SearchThis class
        new SearchThis();
    }
}