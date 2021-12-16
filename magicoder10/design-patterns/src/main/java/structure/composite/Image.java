package structure.composite;

public record Image(String url) implements Component {

    @Override
    public void draw(Graphics graphics) {
        graphics.printf("[Image url=%s]\n", url);
    }
}
