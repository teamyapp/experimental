package structure.composite;

import java.util.List;

public record Container(List<Component> children) implements Component {

    @Override
    public void draw(Graphics graphics) {
        graphics.printf("[Container\n");
        for (Component child : children) {
            graphics.increaseIndentation();
            child.draw(graphics);
            graphics.decreaseIndentation();
        }
        graphics.printf("]\n");
    }
}
