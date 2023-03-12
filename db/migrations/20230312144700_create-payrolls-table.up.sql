begin;

create table payrolls
(
    id        uuid not null,
    tenant_id uuid not null,
    payday    date not null
);
alter table payrolls
    add constraint payrolls_pkey primary key (id);

create table payslips
(
    id         uuid not null,
    tenant_id  uuid not null,
    payroll_id uuid not null,

    gross_pay  int  not null default 0,
    tax        int  not null default 0,
    net_pay    int  not null default 0
);
alter table payslips
    add constraint payslis_pkey primary key (id);
alter table payslips
    add constraint payslips_payrolls_fkey
    foreign key (payroll_id) references payrolls(id);

commit;