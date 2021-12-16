package creation.abstract_factory;

public class SilverWareSet implements FlatWareSet {
    @Override
    public Fork makeFork() {
        return new SilverFork();
    }

    @Override
    public Spoon makeSpoon() {
        return new SilverSpoon();
    }

    @Override
    public Knife makeKnife() {
        return new SilverKnife();
    }
}
