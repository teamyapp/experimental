package creation.prototype;

public class Triangle extends Shape {
    private final int edge1;
    private final int edge2;
    private final int edge3;

    public Triangle(String backgroundColor, int edge1, int edge2, int edge3) {
        super(backgroundColor);
        this.edge1 = edge1;
        this.edge2 = edge2;
        this.edge3 = edge3;
    }

    protected Triangle(Triangle source) {
        super(source);
        this.edge1 = source.edge1;
        this.edge2 = source.edge2;
        this.edge3 = source.edge3;
    }

    @Override
    public Shape clone() {
        return new Triangle(this);
    }

    @Override
    public String toString() {
        return "Triangle{" +
                "edge1=" + edge1 +
                ", edge2=" + edge2 +
                ", edge3=" + edge3 +
                '}';
    }
}
