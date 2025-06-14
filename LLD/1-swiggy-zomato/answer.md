# System requirementfor swiggy-zomato:


## Flow of swiggy-zomato:
user-place-some-order(OrderManagerService)   ->      [ restraunt receive the order(restruantManagerService), delivery agent get notifincation(deliveryManager)]


/// Entities/
class User{
    int userId;
    String name;
    String number;
    Location address;
}

class Items{
    int itemId;
    String itemName;
    double price;
    boolean isAvailable;
}

class Location{
    int locationId;
    String locationName;
}

class RestrantManager{
    int restrauntId;
    String restrauntName;
    List<Items> menuItems;
    Location location;
    boolean isAvailable;
}

enum PaymentMethods{
    UPI, CARD, NETBANKING
}
enum OrderStatus{
    InProcess(1Min Timespan once payment done), Cancel
}

class orderManager{
    int orderid;
    int restrauntId;
    Location userLocation;
    List<Item> orderedItems;
    int userIdOfOrderedBy;
    PaymentMethods paymentMethods;
    boolean orderStatus;
    Location orderedLocation;
    boolean orderStatuas;
    OrderStatus orderStatus;
}
// searching of delivery are done with the range of restraunt, for example in the first try it start with 2-5 km, then extends.

class DeliveryAgent{
    int agentId;
    String name;
    boolean isAvailable;
    Location currentLocation;
}

class DeliveryAssignmentManger{
    int deliveryId;
    int orderId;
    int agentId;
    Location restrauntLocation;
    Location deliveryLocation;
}



// Services:
class OrderManagerService{
    public Order placeOrder(User user, Restraunt resrtaunt, List<Items> orderItems, PaymentMode paymentMode){
        //logic to place order, update payment status, notify restraunt, delivery agent
    }
}

class RestrauntManagerService{
    Public Order receiverOrder(Order order){
        //notify restraunt
    }
}

class DeliveryManagerService{
    public void assignDeliveryAgent(Order order){
        //find nearest available agent and assign. 
    }
}