package creation.abstract_factory;

public abstract class Fork implements Drawable {
    private String heldMeatId;

    public void holdMeat(String meatId) {
        this.heldMeatId = meatId;
    }

    public String releaseMeat() {
        return heldMeatId;
    }
}
