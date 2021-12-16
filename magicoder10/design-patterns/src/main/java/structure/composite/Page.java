package structure.composite;

public class Page implements Component {
    private final String title;
    private final Component child;

    Page(String title, Component child) {
        this.title = title;
        this.child = child;
    }

    @Override
    public void draw(Graphics graphics) {
        graphics.printf("[Page %s\n", title);
        if (child != null) {
            graphics.increaseIndentation();
            child.draw(graphics);
            graphics.decreaseIndentation();
        }
        graphics.printf("]");
    }
}
