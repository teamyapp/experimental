package creation.builder;

public interface DeviceBuilder {
    DeviceBuilder includeBlueTooth();

    DeviceBuilder includeWifi();

    DeviceBuilder includeCamera();

    DeviceBuilder includeSpeaker();

    DeviceBuilder includeMicrophone();

    DeviceBuilder includeGyroscope();

    DeviceBuilder includeAlexaModule();

    DeviceBuilder includeGoogleHomeModule();

    DeviceBuilder reset();
}
