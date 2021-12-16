package creation.prototype;

import java.util.HashMap;
import java.util.Map;

public class ShapeRegistry {
    private final Map<String, Shape> cache = new HashMap<>();

    public static void main(String[] args) {
        ShapeRegistry shapeRegistry = new ShapeRegistry();
        shapeRegistry.put("small-equal-triangle", new Triangle("white", 2, 2, 2));
        shapeRegistry.put("small-square", new Square("white", 2));
        shapeRegistry.put("large-square", new Square("white", 6));

        Shape smallSquare = shapeRegistry.get("small-square");
        System.out.println(smallSquare.toString());

        Shape smallEqualTriangle = shapeRegistry.get("small-equal-triangle");
        System.out.println(smallEqualTriangle.toString());

        Shape largeSquare = shapeRegistry.get("large-square");
        System.out.println(largeSquare.toString());
    }

    public void put(String id, Shape shape) {
        cache.put(id, shape);
    }

    public Shape get(String id) {
        return cache.get(id);
    }
}
