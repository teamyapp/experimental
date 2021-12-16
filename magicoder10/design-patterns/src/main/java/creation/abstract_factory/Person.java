package creation.abstract_factory;

public class Person {
    public static void main(String[] args) {
        haveMeal(new GoldWareSet());
        haveMeal(new SilverWareSet());
    }

    private static void haveMeal(FlatWareSet flatWareSet) {
        Fork fork = flatWareSet.makeFork();
        fork.draw();

        Knife knife = flatWareSet.makeKnife();
        knife.draw();

        Spoon spoon = flatWareSet.makeSpoon();
        spoon.draw();

        fork.holdMeat("meat1");
        knife.cutMeat("meat1");
        spoon.holdSoup("soup1");
        spoon.dropSoup();
    }
}
