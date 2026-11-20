(import '[java.util ArrayList HashMap])

;; Call two Java static methods.
(println "Absolute value:" (Math/abs (long -42)))
(println "Current time in milliseconds:" (System/currentTimeMillis))

(let [clojure-items ["apple" "banana"]
      clojure-stock {"apple" 3, "banana" 5}

      ;; Create Java objects from Clojure collections.
      java-items (ArrayList. clojure-items)
      java-stock (HashMap. clojure-stock)]

  ;; Call Java instance methods.
  (.add java-items "orange")
  (.put java-stock "orange" 2)
  (println "First item:" (.get java-items 0))
  (println "Orange stock:" (.get java-stock "orange"))

  ;; Convert the modified Java collections back to Clojure.
  (println "Clojure vector:" (pr-str (vec java-items)))
  (println "Clojure map:" (pr-str (into {} java-stock))))