class Bug {
    // bug superclass
    String color;
    Integer legs;
    Boolean poisonous;
    void Bite() {
        if (poisonous) {
            System.out.println("You have been poisoned!");
        } else {
            System.out.println("You have been biten!");
        }
    }
}

class Spider extends Bug {
    // spider subclass of bug superclass
    Spider() {
        // spider constructor
        legs = 8;
        color = "Black";
        poisonous = true;
    }
}

class Roach extends Bug {
    // roach subclass of bug superclass
    Roach() {
        // roach constructor
        legs = 6;
        color = "Brown";
        poisonous = false;
    }
}

public class Main {
    public static void main(String[] args) {
        Bug blackwidow = new Spider();
        blackwidow.Bite();
    }
}