
(println "-- DRONE DELIVERY SYSTEM --")

;; 1. LIST - Delivery locations
(def locations (list "Rohini" "Dwarka" "Janakpuri"))

(println "\nDelivery Locations:" locations)
(println "First Location:" (first locations))
(println "Other Locations:" (rest locations))

(def updated-locations (conj locations "Pitampura"))

(println "After Adding Location:" updated-locations)
(println "Original Locations:" locations)


;; 2. VECTOR - Drone models
(def drones ["SkyLite" "AirPro" "Falcon"])

(println "\nAvailable Drones:" drones)
(println "Drone at Index 1:" (get drones 1))
(println "Number of Drones:" (count drones))

(def new-drones (conj drones "AeroX"))

(println "After Adding Drone:" new-drones)
(println "Original Drone List:" drones)


;; 3. MAP - Delivery details
(def deliveries
  {"D101" "Rohini"
   "D102" "Dwarka"
   "D103" "Janakpuri"})

(println "\nDelivery Records:" deliveries)
(println "D102 Destination:" (get deliveries "D102"))

(def added-delivery (assoc deliveries "D104" "Pitampura"))

(println "After New Booking:" added-delivery)

(def removed-delivery (dissoc added-delivery "D101"))

(println "After Cancelling D101:" removed-delivery)
(println "Original Records:" deliveries)


;; 4. SET - Available delivery services
(def services #{"Food" "Medicine" "Documents"})

(println "\nAvailable Services:" services)
(println "Medicine Service Available:" (contains? services "Medicine"))

(def extra-service (conj services "Electronics"))

(println "After Adding Service:" extra-service)

(def final-services (disj extra-service "Food"))

(println "After Removing Food:" final-services)
(println "Original Services:" services)


(println "\nDrone delivery operations completed.")
