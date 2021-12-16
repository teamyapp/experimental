package creation.singleton;

public final class EventBus {
    // volatile solves cache coherence issue by ensuring writes always occur before reads
    private static volatile EventBus instance;

    private EventBus() {
    }

    public static void main(String[] args) {
        EventBus.GetInstance().publish("Start", "Hello world!");
    }

    public void publish(String event, String message) {
        System.out.printf("[%s] %s\n", event, message);
    }

    public static EventBus GetInstance() {
        if (instance == null) {
            synchronized (EventBus.class) {
                if (instance == null) {
                    instance = new EventBus();
                }
            }
        }
        return instance;
    }
}
