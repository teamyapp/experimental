package creation.abstract_factory;

public class GoldWareSet implements FlatWareSet {
    @Override
    public Fork makeFork() {
        return new GoldFork();
    }

    @Override
    public Spoon makeSpoon() {
        return new GoldSpoon();
    }

    @Override
    public Knife makeKnife() {
        return new GoldKnife();
    }
}
