CREATE TABLE "users"(
    "id" UUID NOT NULL,
    "lastname" VARCHAR(255) NOT NULL,
    "firstname" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "roles" VARCHAR(255) NOT NULL,
    "salon_id" UUID NULL
);
ALTER TABLE
    "users" ADD PRIMARY KEY("id");
ALTER TABLE
    "users" ADD CONSTRAINT "users_email_unique" UNIQUE("email");
CREATE TABLE "reservation"(
    "id" UUID NOT NULL,
    "slot_id" UUID NOT NULL,
    "customer_id" UUID NOT NULL
);
ALTER TABLE
    "reservation" ADD PRIMARY KEY("id");
CREATE TABLE "services"(
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "time" TIME(0) WITHOUT TIME ZONE NOT NULL,
    "price" DOUBLE PRECISION NOT NULL,
    "salon_id" UUID NOT NULL
);
ALTER TABLE
    "services" ADD PRIMARY KEY("id");
CREATE TABLE "salons"(
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "salons" ADD PRIMARY KEY("id");
CREATE TABLE "slots"(
    "id" UUID NOT NULL,
    "date" DATE NOT NULL,
    "beginning_hour" TIME(0) WITHOUT TIME ZONE NOT NULL,
    "end_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
    "hairdressing_staff_id" UUID NOT NULL,
    "salon_id" UUID NOT NULL
);
ALTER TABLE
    "slots" ADD PRIMARY KEY("id");
ALTER TABLE
    "reservation" ADD CONSTRAINT "reservation_slot_id_foreign" FOREIGN KEY("slot_id") REFERENCES "slots"("id");
ALTER TABLE
    "slots" ADD CONSTRAINT "slots_hairdressing_staff_id_foreign" FOREIGN KEY("hairdressing_staff_id") REFERENCES "users"("id");
ALTER TABLE
    "reservation" ADD CONSTRAINT "reservation_customer_id_foreign" FOREIGN KEY("customer_id") REFERENCES "users"("id");
ALTER TABLE
    "slots" ADD CONSTRAINT "slots_salon_id_foreign" FOREIGN KEY("salon_id") REFERENCES "salons"("id");
ALTER TABLE
    "services" ADD CONSTRAINT "services_salon_id_foreign" FOREIGN KEY("salon_id") REFERENCES "salons"("id");
ALTER TABLE
    "users" ADD CONSTRAINT "users_salon_id_foreign" FOREIGN KEY("salon_id") REFERENCES "salons"("id");

INSERT INTO "users" ("id", "lastname", "firstname", "email", "password", "roles", "salon_id")
    VALUES ('ce801349-752c-4cee-a660-9f7cceaf7131', 'admin', 'admin', 'admin', '$2a$10$kDjL2RObjxHJbue3uX9UzeZxcmxhE2ig0wZWgzsXZk4hXdVB9gHRm', 'admin', NULL);