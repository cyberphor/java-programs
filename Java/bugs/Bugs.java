class Spider {
    Integer legs = 8;
}

class Roach {
    Integer legs = 6;
}

public class Bugs {
    public static void main(String[] args) {
        Spider spider = new Spider();
        Roach roach = new Roach();
        System.out.println("Spiders have " + spider.legs + " legs.");
        System.out.println("Roaches have " + roach.legs + " legs.");
    }
}