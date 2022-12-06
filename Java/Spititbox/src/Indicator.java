import java.time.LocalDateTime;

public class Indicator {
    private LocalDateTime FirstSeen = LocalDateTime.now();
    public void getFirstSeen() {
        System.out.println(this.FirstSeen);
    }

    public void setFirstSeen() {
        // code goes here;
    }
}
