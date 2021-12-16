package structure.composite;

public record Button(String text) implements Component {

    @Override
    public void draw(Graphics graphics) {
        graphics.printf("[Button text=%s]\n", text);
    }
}
