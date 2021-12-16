package creation.abstract_factory;

public abstract class Spoon implements Drawable {
    private String heldSoupId;

    public void holdSoup(String soupId) {
        this.heldSoupId = soupId;
    }

    public String dropSoup() {
        return heldSoupId;
    }
}
