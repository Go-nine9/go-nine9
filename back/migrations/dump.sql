CREATE TABLE "users"(
    "id" UUID NOT NULL,
    "lastname" VARCHAR(255) NOT NULL,
    "firstname" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "roles" VARCHAR(255) NOT NULL,
    "salon_id" UUID NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);
ALTER TABLE
    "users" ADD PRIMARY KEY("id");
ALTER TABLE
    "users" ADD CONSTRAINT "users_email_unique" UNIQUE("email");
CREATE TABLE "reservations"(
    "id" UUID NOT NULL,
    "slot_id" UUID NOT NULL,
    "customer_id" UUID NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);
ALTER TABLE
    "reservations" ADD PRIMARY KEY("id");
CREATE TABLE "services"(
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "time" TIME(0) WITHOUT TIME ZONE NOT NULL,
    "price" DOUBLE PRECISION NOT NULL,
    "salon_id" UUID NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);
ALTER TABLE
    "services" ADD PRIMARY KEY("id");
CREATE TABLE "salons"(
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);
ALTER TABLE
    "salons" ADD PRIMARY KEY("id");
CREATE TABLE "slots"(
    "id" UUID NOT NULL,
    "date" DATE NOT NULL,
    "beginning_hour" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "end_time" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "hairdressing_staff_id" UUID NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);
ALTER TABLE
    "slots" ADD PRIMARY KEY("id");
ALTER TABLE
    "reservations" ADD CONSTRAINT "reservations_slot_id_foreign" FOREIGN KEY("slot_id") REFERENCES "slots"("id");
ALTER TABLE
    "slots" ADD CONSTRAINT "slots_hairdressing_staff_id_foreign" FOREIGN KEY("hairdressing_staff_id") REFERENCES "users"("id");
ALTER TABLE
    "reservations" ADD CONSTRAINT "reservations_customer_id_foreign" FOREIGN KEY("customer_id") REFERENCES "users"("id");
ALTER TABLE
    "services" ADD CONSTRAINT "services_salon_id_foreign" FOREIGN KEY("salon_id") REFERENCES "salons"("id");
ALTER TABLE
    "users" ADD CONSTRAINT "users_salon_id_foreign" FOREIGN KEY("salon_id") REFERENCES "salons"("id");

INSERT INTO "salons" ("id", "name", "address", "phone")
    VALUES ('ce801349-752c-4cee-a660-9f7cceaf7132', 'Hairstyle', '18 rue Paris,75016', '010230');

INSERT INTO "users" ("id", "lastname", "firstname", "email", "password", "roles", "salon_id")
    VALUES ('9f7ceeaf-7132-4cee-a660-ce801349-852c', 'Lastname', 'Firstname', 'email@example.com', 'password', 'admin', 'ce801349-752c-4cee-a660-9f7cceaf7132');

INSERT INTO "users" ("id", "lastname", "firstname", "email", "password", "roles", "salon_id")
    VALUES ('9f7ceeaf-7132-4cee-a660-ce801349-853c', 'admin', 'admin', 'admin', 'admin', 'admin', NULL);

INSERT INTO "slots"("id","date","beginning_hour","end_time","hairdressing_staff_id")
    VALUES ('9f7ceeaf-7132-4cee-a660-ce801349-854c','2021-01-01','2024-01-21 22:57:35','2024-01-21 23:57:35','9f7ceeaf-7132-4cee-a660-ce801349-852c');

INSERT INTO "reservations"("id","slot_id","customer_id")
    VALUES ('9f7ceeaf-7132-4cee-a660-ce801349-855c', '9f7ceeaf-7132-4cee-a660-ce801349-854c', '9f7ceeaf-7132-4cee-a660-ce801349-853c');
