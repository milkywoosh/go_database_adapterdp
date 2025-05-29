# Container name
CONTAINER_NAME = pg17v1
EXEC_DOCKER_PG17V1 = docker exec -it pg17v1 psql -U postgres -d toko_buku_online_nextjs -c

# Default target
.PHONY: all
all: help


# Display help
.PHONY: help
help:
	@echo "Makefile for managing PostgreSQL Docker container $(CONTAINER_NAME):"
	@echo "  make start   - Start the container"
	@echo "  make stop    - Stop the container"
	@echo "  make restart - Restart the container"
	@echo "  make status  - Check container status"
	@echo "  make help    - Show this help message"

.PHONY: start
start:
	@echo "Starting container $(CONTAINER_NAME)..."
	docker start $(CONTAINER_NAME)

.PHONY: stop
stop:
	@echo "Stop container $(CONTAINER_NAME)"
	docker stop $(CONTAINER_NAME)

.PHONY: status
status:
	@echo "Checking status of container $(CONTAINER_NAME)..."
	docker ps -a --filter "name=$(CONTAINER_NAME)"

.PHONY: toko_buku_users
toko_buku_users:
	$(EXEC_DOCKER_PG17V1)\
	"select * from users limit 5"

.PHONY: check_user_roles
check_user_roles:
	$(EXEC_DOCKER_PG17V1)\
	"select u.username, r.role_name from users u \
	left join user_roles ur on u.id = ur.user_id \
	left join roles r on r.id = ur.role_id \
	where username='luke'"

.PHONY: run_any_query
run_any_query:
	$(EXEC_DOCKER_PG17V1)\
	"select * \
	from users u"

.PHONY: find_username
find_username:
	$(EXEC_DOCKER_PG17V1)\
	"select u.username \
	from users u \
	where u.username='luke'"

.PHONY: check_roles
check_roles:
	$(EXEC_DOCKER_PG17V1)\
	"SELECT r.role_name, u.username from user_roles ur \
	INNER JOIN users u on u.id = ur.user_id \
	INNER JOIN roles r on r.id = ur.role_id \
	where u.username = 'ben'"

# PURCHASE NUMBER PURCHASE_NUMBER
.PHONY: check_purchase_number
check_purchase_number:
	$(EXEC_DOCKER_PG17V1)\
	"SELECT * FROM purchase_items pn \
	WHERE pn.purchase_number = 'PRCBOOK2025531955112124'"

.PHONY: check_purchase_history
check_purchase_history:
	$(EXEC_DOCKER_PG17V1)\
	"SELECT * FROM purchase_histories ph \
	WHERE ph.purchase_number = 'PRCBOOK2025731955112345'"

.PHONY: trx_create_prc
trx_create_prc:
	$(EXEC_DOCKER_PG17V1)\
	"INSERT INTO purchase_histories \
		(date_of_sale, customer_id, total_price_payment, status, purchase_number) \
		VALUES('2024-12-17 23:20:17.925', 4, 100000.000, 'pending', 'PRCBOOK2025731955112345') \
	"

.PHONY: datatable_prc
datatable_prc:
	$(EXEC_DOCKER_PG17V1) \
	"select  \
		pi.purchase_number as purchase_number, \
		pi.total_price, \
		pi.qty, \
		b.title, \
		b.price as price_each, \
		ph.date_of_sale, \
		ph.status \
	from purchase_items pi \
	left join books b on b.id = pi.book_id  \
	left join purchase_histories ph on ph.id = pi.purchase_history_id \
	where pi.purchase_number IS NOT NULL \
	"