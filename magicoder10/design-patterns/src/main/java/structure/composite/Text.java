package structure.composite;

public record Text(String text) implements Component {

    @Override
    public void draw(Graphics graphics) {
        graphics.printf("[Text text=%s]\n", text);
    }
}
