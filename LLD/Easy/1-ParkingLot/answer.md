# Parking LOT Requirement: 
        =>
        1. The parking lot should have multiple levels, each level with a certain number of parking spots.
        2. The parking lot should support different types of vehicles, such as cars, motorcycles, and trucks
        3. Each parking spot should be able to accomodate a specific type of vehicle.
        4. The system should assign a parking spot to a vehicle upon entry and release it when the vehicle exits.
        5. The system should track the availability of parking spots and provide real-time information to customers.
        6. The system should handle multiple entry and exit points and support concurrent access.


# Flow of the application:
        => System will automatically assign the available parking slot to vehicle.

class ParkingSlots{
    int slotId;
    boolean isAvailable;
}
class parkingLevel{
    int level;
    int capacity;
    List<ParkingSlots> totalAvailable;
}
enum VehiclesTypes{
    CARS, MOTORCYCLES, TRUCKS
}
class Vehicle{
    String licence;
    VehicleTypes vehicles;
}
class ParkingManagement{
    int parkingId;
    user user
}