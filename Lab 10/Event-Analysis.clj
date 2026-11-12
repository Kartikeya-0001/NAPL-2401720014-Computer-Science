
(println "-- COLLEGE EVENT ANALYSIS --")

;; Expenses for different event activities
(def event-costs [850 1200 450 1750 950 600])

(println "\nEvent Expenses:" event-costs)

;; map - Add 100 to each expense for extra charges
(def revised-costs
  (map #(+ % 100) event-costs))

(println "Revised Expenses:" revised-costs)

;; filter - Find expenses above 1000
(def costly-activities
  (filter #(> % 1000) revised-costs))

(println "Expenses Above 1000:" costly-activities)

;; reduce - Calculate total event budget
(def total-budget
  (reduce + revised-costs))

(println "Total Event Budget:" total-budget)

;; take and drop
(println "First Three Expenses:" (take 3 event-costs))
(println "Remaining Expenses:" (drop 3 event-costs))


;; Second collection - Tickets sold for different events
(def tickets-sold '(18 32 12 45 27 15))

(println "\nTickets Sold:" tickets-sold)

;; map - Add 5 more tickets to each event
(def revised-tickets
  (map #(+ % 5) tickets-sold))

(println "Updated Ticket Counts:" revised-tickets)

;; filter - Find events with more than 30 tickets
(def popular-events
  (filter #(> % 30) revised-tickets))

(println "Popular Events:" popular-events)

;; reduce - Calculate total tickets
(def total-tickets
  (reduce + revised-tickets))

(println "Total Tickets:" total-tickets)

;; take and drop
(println "First Two Events:" (take 2 tickets-sold))
(println "Other Events:" (drop 2 tickets-sold))


;; Threading macro - Select expensive activities
(def selected-costs
  (->> event-costs
       (map #(+ % 100))
       (filter #(> % 1000))
       (take 2)))

(println "\nSelected Expenses Using ->>:" selected-costs)

(println "\nEvent analysis completed.")
