# Bare Metal Orchestrator

This is the backend code for the Bare Metal Orchestrator (BMO) - it includes everything needed to launch the project. The `bmo-ipmi-firmware` is also included as a submodule - this is the code that runs on the Arduino.

## What You Need

- **Raspberry Pi** (or any low power SBC) to act as the main orchestrator. The Pi runs the API server and sends commands to the Arduino.
- **Arduino Nano 33 IoT** to interface with the server hardware.
- **Compatible Motherboard** with connections for the `PWRBTN`, `RESET`, and `PLED` pins.
- **Docker & Docker Compose** installed on the Raspberry Pi.

## How to Run

To get started, simply run the following command inside this directory:

```bash
docker-compose up --build -d
```

And that’s it.

Currently, the project is hardcoded to use `/dev/ttyACM0` as the serial port. I’m working on an update to automatically find the correct serial port without needing it hardcoded. I might also experiment with using the Raspberry Pi’s GPIO pins directly - bypassing the need for USB communication with the Arduino.


**Note**: This project is still a work in progress (WIP) and will be updated frequently. Stay tuned!