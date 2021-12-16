package creation.builder;

public class Assembler {
    public static void makeAlexaRobotKit(DeviceBuilder builder) {
        builder.reset()
                .includeSpeaker()
                .includeMicrophone()
                .includeWifi()
                .includeBlueTooth()
                .includeAlexaModule();
    }

    public static void makeGoogleRobotKit(DeviceBuilder builder) {
        builder.reset()
                .includeSpeaker()
                .includeMicrophone()
                .includeWifi()
                .includeBlueTooth()
                .includeGoogleHomeModule();
    }

    public static void main(String[] args) {
        RoboticKit.Builder robotBuilder = new RoboticKit.Builder();
        RemoteControl.Builder remoteControlBuilder = new RemoteControl.Builder();

        Assembler.makeSimpleRobotKit(robotBuilder);
        Assembler.makeSimpleRobotKit(remoteControlBuilder);

        RoboticKit roboticKit = robotBuilder.Build();
        RemoteControl remoteControl = remoteControlBuilder.Build();

        remoteControl.pairWith(roboticKit);
        remoteControl.tryPlaySpotify();

        Assembler.makePremiumRobotKit(robotBuilder);
        Assembler.makePremiumRobotKit(remoteControlBuilder);

        roboticKit = robotBuilder.Build();
        remoteControl = remoteControlBuilder.Build();
        remoteControl.pairWith(roboticKit);
        remoteControl.tryPlaySpotify();
    }

    public static void makeSimpleRobotKit(DeviceBuilder builder) {
        builder.reset()
                .includeSpeaker()
                .includeMicrophone()
                .includeWifi()
                .includeBlueTooth();
    }

    public static void makePremiumRobotKit(DeviceBuilder builder) {
        builder.reset()
                .includeSpeaker()
                .includeMicrophone()
                .includeWifi()
                .includeBlueTooth()
                .includeCamera()
                .includeGyroscope()
                .includeAlexaModule()
                .includeGoogleHomeModule();
    }
}
