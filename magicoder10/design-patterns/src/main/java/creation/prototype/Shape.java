package creation.prototype;

public abstract class Shape {
    private final String backgroundColor;

    public Shape(String backgroundColor) {
        this.backgroundColor = backgroundColor;
    }

    protected Shape(Shape source) {
        this.backgroundColor = source.backgroundColor;
    }

    public abstract Shape clone();

    @Override
    public String toString() {
        return "Shape{" +
                "backgroundColor='" + backgroundColor + '\'' +
                '}';
    }
}
