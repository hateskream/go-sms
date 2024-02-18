CREATE TABLE "public"."cars" (
    "id" int4 NOT NULL DEFAULT nextval('cars_id_seq'::regclass),
    "number" varchar DEFAULT 'NULL'::character varying,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."features" (
    "id" int4 NOT NULL DEFAULT nextval('features_id_seq'::regclass),
    "name" varchar NOT NULL DEFAULT ''::character varying,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."pricing_groups" (
    "id" int4 NOT NULL DEFAULT nextval('space_groups_id_seq'::regclass),
    "name" varchar NOT NULL DEFAULT 'EMPTY'::character varying,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."reservation_statuses" (
    "id" int4 NOT NULL DEFAULT nextval('reservation_statuses_id_seq'::regclass),
    "name" varchar NOT NULL DEFAULT '''active'''::character varying,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."space_features" (
    "id" int4 NOT NULL DEFAULT nextval('untitled_table_214_id_seq'::regclass),
    "space_id" int4,
    "feature_id" int4,
    "is_required" bool NOT NULL DEFAULT false,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."space_occupancy" (
    "id" int4 NOT NULL DEFAULT nextval('space_occupancy_id_seq'::regclass),
    "space_id" int4,
    "timestamp" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "is_occupied" bool NOT NULL DEFAULT false,
    "car_id" int4,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."space_reservations" (
    "id" int4 NOT NULL DEFAULT nextval('space_reservetions_id_seq'::regclass),
    "time_from" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "time_to" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "car_id" int4,
    "reservation_fee" float4 NOT NULL DEFAULT 0,
    "status_id" int4,
    "space_id" int4,
    "parking_time_from" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "parking_time_to" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "parking_fee" float4 NOT NULL DEFAULT 0,
    "parking_fee_breakdown" float4 NOT NULL DEFAULT 0,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."space_statuses" (
    "id" int4 NOT NULL DEFAULT nextval('space_statuses_id_seq'::regclass),
    "name" varchar NOT NULL DEFAULT 'EMPTY'::character varying,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."spaces" (
    "id" int4 NOT NULL DEFAULT nextval('untitled_table_210_id_seq'::regclass),
    "name" varchar NOT NULL DEFAULT 'EMPTY'::character varying,
    "physical_id" int4,
    "group_id" int4,
    "status_id" int4,
    "has_camera" bool NOT NULL DEFAULT false,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."time_pricing_policy" (
    "id" int4 NOT NULL DEFAULT nextval('time_pricing_policy_id_seq'::regclass),
    "rate" float4 NOT NULL DEFAULT 0,
    "hour" int2 NOT NULL DEFAULT 0,
    "day_of_week" int2 NOT NULL DEFAULT 0,
    "group_id" int4 NOT NULL DEFAULT 1,
    PRIMARY KEY ("id")
);

ALTER TABLE "public"."space_features" ADD FOREIGN KEY ("feature_id") REFERENCES "public"."features"("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "public"."space_features" ADD FOREIGN KEY ("space_id") REFERENCES "public"."spaces"("id");
ALTER TABLE "public"."space_occupancy" ADD FOREIGN KEY ("car_id") REFERENCES "public"."cars"("id");
ALTER TABLE "public"."space_occupancy" ADD FOREIGN KEY ("space_id") REFERENCES "public"."spaces"("id");
ALTER TABLE "public"."space_reservations" ADD FOREIGN KEY ("space_id") REFERENCES "public"."spaces"("id");
ALTER TABLE "public"."space_reservations" ADD FOREIGN KEY ("status_id") REFERENCES "public"."reservation_statuses"("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "public"."space_reservations" ADD FOREIGN KEY ("car_id") REFERENCES "public"."cars"("id");
ALTER TABLE "public"."spaces" ADD FOREIGN KEY ("status_id") REFERENCES "public"."space_statuses"("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "public"."spaces" ADD FOREIGN KEY ("group_id") REFERENCES "public"."pricing_groups"("id");
ALTER TABLE "public"."time_pricing_policy" ADD FOREIGN KEY ("group_id") REFERENCES "public"."pricing_groups"("id") ON UPDATE CASCADE;
