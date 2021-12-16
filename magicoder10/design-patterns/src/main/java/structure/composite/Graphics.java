package structure.composite;

public class Graphics {
    private final StringBuilder stringBuilder = new StringBuilder();
    private int indent;

    void increaseIndentation() {
        indent++;
    }

    void decreaseIndentation() {
        indent--;
    }

    void printf(String format, Object... args) {
        System.out.print(" ".repeat(indent * 2));
        System.out.printf(format, args);
    }
}
