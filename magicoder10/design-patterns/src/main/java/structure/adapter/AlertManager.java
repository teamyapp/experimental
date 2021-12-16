package structure.adapter;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public record AlertManager(List<Notifier> notifiers) {
    public AlertManager {
        if (notifiers == null) {
            notifiers = new ArrayList<>();
        }
    }

    public static void main(String[] args) {
        AlertManager alertManager = new AlertManager(
                Arrays.asList(new SlackNotifierAdapter("api_key", "outage"))
        );

        alertManager.reportOutage("OMG!");
    }

    void reportOutage(String errMessage) {
        for (Notifier notifier : notifiers) {
            notifier.notify(errMessage);
        }
    }
}
