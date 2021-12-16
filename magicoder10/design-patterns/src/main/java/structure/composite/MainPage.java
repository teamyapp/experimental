package structure.composite;

import java.util.Arrays;

public class MainPage extends Page {
    public MainPage() {
        super("Main", new Container(Arrays.asList(
                new Container(Arrays.asList(
                        new Button("btn 0"),
                        new Image("https://www.teamyapp.com")
                )),
                new Text("Welcome!"))
        ));
    }

    public static void main(String[] args) {
        Graphics graphics = new Graphics();
        MainPage mainPage = new MainPage();
        mainPage.draw(graphics);

    }
}
