package creation.abstract_factory;

import java.util.ArrayList;
import java.util.List;

public abstract class Knife implements Drawable {
    public List<String> cutMeat(String meatId) {
        return new ArrayList<>();
    }
}
