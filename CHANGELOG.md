v 1.0.4
  - debug: rename -d flag to -debug. #38
  - Remove hacks for pay_amount and payment_id with custom unmarshalling. #36

v 1.0.3
  - Unmarshal error on payment creation in production environment. #34

v 1.0.2
  - sandbox: add support for optional case for new payments. #32
  - payment: missing PayAmount and PayCurrency fields. #30

v 1.0.1
  - Explicitely add mocks directory and remove pkg package. #28
  - Add Go report card to README. #24
  - Link methods to Go documentation. #22

v 1.0.0
  - cmd/np: add -f flag to read config file. #20
  - Implement create payment from invoice. #15
  - Implement create invoice. #14
  - Implement get/update payment estimate. #16
  - Implement get available checked currencies. #12
  - Implement get list of payments. #10
  - Implement create payment. #8
  - Add more unit tests. #6
  - Implement a couple of endpoints. #2
  - Setup testing and mocking service structure. #3
