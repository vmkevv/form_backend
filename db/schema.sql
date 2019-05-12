create table user(
id serial primary key,
nombre varchar (127) not null,
appat varchar (127) not null,
apmat varchar (127) not null ,
ci varchar(127) not null,
cel int,
correo varchar (127) not null,
dir varchar (255) not null,
createdAt timestamp
);
