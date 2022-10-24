public class Add {
    public static void main(String[] args){
        if (args.length == 2) {
            Integer solution = Integer.parseInt(args[0]) + Integer.parseInt(args[1]);
            System.out.println("Problem: " + args[0] + " + " + args[1]);
            System.out.println("Solution: " + solution);
        }
    }
}