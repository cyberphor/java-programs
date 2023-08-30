public class Calculator {
    static void add(String[] args) {
        if (args.length == 2) {
            Integer x = Integer.parseInt(args[0]);
            Integer y = Integer.parseInt(args[1]);
            Integer solution = x + y;
            System.out.println("Problem: " + x + " + " + y);
            System.out.println("Solution: " + solution);
        }
    }

    public static void main(String[] args) {
        add(args);
    }
}