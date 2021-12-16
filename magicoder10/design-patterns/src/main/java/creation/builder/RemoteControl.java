package creation.builder;

public class RemoteControl {
    private boolean hasBlueTooth;
    private boolean hasWifi;
    private boolean hasCamera;
    private boolean hasSpeaker;
    private boolean hasMicrophone;
    private boolean hasGyroscope;
    private boolean supportAlexa;
    private boolean supportGoogleHome;

    private RoboticKit roboticKit;

    private RemoteControl(
            boolean hasBlueTooth,
            boolean hasWifi,
            boolean hasCamera,
            boolean hasSpeaker,
            boolean hasMicrophone,
            boolean hasGyroscope,
            boolean supportAlexa,
            boolean supportGoogleHome
    ) {
        this.hasBlueTooth = hasBlueTooth;
        this.hasWifi = hasWifi;
        this.hasCamera = hasCamera;
        this.hasSpeaker = hasSpeaker;
        this.hasMicrophone = hasMicrophone;
        this.hasGyroscope = hasGyroscope;
        this.supportAlexa = supportAlexa;
        this.supportGoogleHome = supportGoogleHome;
    }

    public void pairWith(RoboticKit roboticKit) {
        this.roboticKit = roboticKit;
    }

    public void tryPlaySpotify() {
        roboticKit.tryPlaySpotify();
    }

    public static class Builder implements DeviceBuilder {
        private boolean hasBlueTooth;
        private boolean hasWifi;
        private boolean hasCamera;
        private boolean hasSpeaker;
        private boolean hasMicrophone;
        private boolean hasGyroscope;
        private boolean supportAlexa;
        private boolean supportGoogleHome;

        @Override
        public Builder includeBlueTooth() {
            hasBlueTooth = true;
            return this;
        }

        @Override
        public Builder includeWifi() {
            hasWifi = true;
            return this;
        }

        @Override
        public Builder includeCamera() {
            hasCamera = true;
            return this;
        }

        @Override
        public Builder includeSpeaker() {
            hasSpeaker = true;
            return this;
        }

        @Override
        public Builder includeMicrophone() {
            hasMicrophone = true;
            return this;
        }

        @Override
        public Builder includeGyroscope() {
            hasGyroscope = true;
            return this;
        }

        @Override
        public Builder includeAlexaModule() {
            supportAlexa = true;
            return this;
        }

        @Override
        public Builder includeGoogleHomeModule() {
            supportGoogleHome = true;
            return this;
        }

        @Override
        public DeviceBuilder reset() {
            this.hasBlueTooth = false;
            this.hasWifi = false;
            this.hasCamera = false;
            this.hasSpeaker = false;
            this.hasMicrophone = false;
            this.hasGyroscope = false;
            this.supportAlexa = false;
            this.supportGoogleHome = false;
            return this;
        }

        public RemoteControl Build() {
            return new RemoteControl(
                    hasBlueTooth,
                    hasWifi,
                    hasCamera,
                    hasSpeaker,
                    hasMicrophone,
                    hasGyroscope,
                    supportAlexa,
                    supportGoogleHome
            );
        }
    }
}
