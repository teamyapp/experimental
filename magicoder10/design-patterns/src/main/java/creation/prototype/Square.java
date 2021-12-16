package creation.prototype;

public class Square extends Shape {
    private final int width;

    public Square(String backgroundColor, int width) {
        super(backgroundColor);
        this.width = width;
    }

    protected Square(Square source) {
        super(source);
        this.width = source.width;
    }

    @Override
    public Square clone() {
        return new Square(this);
    }

    @Override
    public String toString() {
        return "Square{" +
                "width=" + width +
                '}';
    }
}
