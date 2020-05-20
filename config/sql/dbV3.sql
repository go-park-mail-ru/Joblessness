--
-- PostgreSQL database dump
--

-- Dumped from database version 10.12 (Ubuntu 10.12-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.12 (Ubuntu 10.12-0ubuntu0.18.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: education; Type: TABLE; Schema: public; Owner: huvalk
--

CREATE TABLE public.education (
    summary_id integer NOT NULL,
    institution character varying(60),
    speciality character varying(60),
    graduated timestamp without time zone,
    type character varying(60)
);


ALTER TABLE public.education OWNER TO huvalk;

--
-- Name: experience; Type: TABLE; Schema: public; Owner: huvalk
--

CREATE TABLE public.experience (
    company_name character varying(60),
    role character varying(120),
    responsibilities text,
    start timestamp without time zone,
    stop timestamp without time zone,
    summary_id integer NOT NULL
);


ALTER TABLE public.experience OWNER TO huvalk;

--
-- Name: favorite; Type: TABLE; Schema: public; Owner: huvalk
--

CREATE TABLE public.favorite (
    user_id integer NOT NULL,
    favorite_id integer NOT NULL
);


ALTER TABLE public.favorite OWNER TO huvalk;

--
-- Name: organization; Type: TABLE; Schema: public; Owner: huvalk
--

CREATE TABLE public.organization (
    id integer NOT NULL,
    name character varying(60) NOT NULL,
    site character varying(60),
    about text DEFAULT ''::text
);


ALTER TABLE public.organization OWNER TO huvalk;

--
-- Name: organization_id_seq; Type: SEQUENCE; Schema: public; Owner: huvalk
--

CREATE SEQUENCE public.organization_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.organization_id_seq OWNER TO huvalk;

--
-- Name: organization_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: huvalk
--

ALTER SEQUENCE public.organization_id_seq OWNED BY public.organization.id;


--
-- Name: person; Type: TABLE; Schema: public; Owner: huvalk
--

CREATE TABLE public.person (
    id integer NOT NULL,
    name character varying(60) NOT NULL,
    gender character varying(10),
    birthday date
);


ALTER TABLE public.person OWNER TO huvalk;

--
-- Name: person_id_seq; Type: SEQUENCE; Schema: public; Owner: huvalk
--

CREATE SEQUENCE public.person_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.person_id_seq OWNER TO huvalk;

--
-- Name: person_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: huvalk
--

ALTER SEQUENCE public.person_id_seq OWNED BY public.person.id;


--
-- Name: requirement; Type: TABLE; Schema: public; Owner: huvalk
--

CREATE TABLE public.requirement (
    summary_id integer,
    driver_license character varying(60),
    has_car boolean DEFAULT false,
    schedule text,
    employment text
);


ALTER TABLE public.requirement OWNER TO huvalk;

--
-- Name: response; Type: TABLE; Schema: public; Owner: huvalk
--

CREATE TABLE public.response (
    summary_id integer,
    vacancy_id integer,
    rejected boolean DEFAULT false,
    approved boolean DEFAULT false,
    date timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT approved_rejected CHECK (((rejected = false) OR (approved = false))),
    CONSTRAINT has_reflection CHECK (((rejected IS NOT NULL) AND (approved IS NOT NULL)))
);


ALTER TABLE public.response OWNER TO huvalk;

--
-- Name: session; Type: TABLE; Schema: public; Owner: huvalk
--

CREATE TABLE public.session (
    user_id integer,
    session_id character varying(64) NOT NULL,
    expires timestamp without time zone
);


ALTER TABLE public.session OWNER TO huvalk;

--
-- Name: summary; Type: TABLE; Schema: public; Owner: huvalk
--

CREATE TABLE public.summary (
    id integer NOT NULL,
    author integer,
    keywords text
);


ALTER TABLE public.summary OWNER TO huvalk;

--
-- Name: summary_id_seq; Type: SEQUENCE; Schema: public; Owner: huvalk
--

CREATE SEQUENCE public.summary_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.summary_id_seq OWNER TO huvalk;

--
-- Name: summary_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: huvalk
--

ALTER SEQUENCE public.summary_id_seq OWNED BY public.summary.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: huvalk
--

CREATE TABLE public.users (
    id integer NOT NULL,
    login character varying(60) NOT NULL,
    password character varying(60) NOT NULL,
    organization_id integer,
    person_id integer,
    tag character varying(20),
    email character varying(60),
    phone character varying(20),
    registered timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    avatar text DEFAULT 'https://hb.bizmrg.com/imgs-hh/default-avatar.png'::text,
    CONSTRAINT entity_not_empty CHECK (((organization_id IS NOT NULL) OR (person_id IS NOT NULL)))
);


ALTER TABLE public.users OWNER TO huvalk;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: huvalk
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO huvalk;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: huvalk
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: vacancy; Type: TABLE; Schema: public; Owner: huvalk
--

CREATE TABLE public.vacancy (
    id integer NOT NULL,
    organization_id integer,
    name character varying(60) NOT NULL,
    description text,
    with_tax boolean DEFAULT false,
    responsibilities text,
    conditions text,
    keywords text,
    salary_from integer DEFAULT 0 NOT NULL,
    salary_to integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.vacancy OWNER TO huvalk;

--
-- Name: vacancy_id_seq; Type: SEQUENCE; Schema: public; Owner: huvalk
--

CREATE SEQUENCE public.vacancy_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vacancy_id_seq OWNER TO huvalk;

--
-- Name: vacancy_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: huvalk
--

ALTER SEQUENCE public.vacancy_id_seq OWNED BY public.vacancy.id;


--
-- Name: organization id; Type: DEFAULT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.organization ALTER COLUMN id SET DEFAULT nextval('public.organization_id_seq'::regclass);


--
-- Name: person id; Type: DEFAULT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.person ALTER COLUMN id SET DEFAULT nextval('public.person_id_seq'::regclass);


--
-- Name: summary id; Type: DEFAULT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.summary ALTER COLUMN id SET DEFAULT nextval('public.summary_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: vacancy id; Type: DEFAULT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.vacancy ALTER COLUMN id SET DEFAULT nextval('public.vacancy_id_seq'::regclass);


--
-- Data for Name: education; Type: TABLE DATA; Schema: public; Owner: huvalk
--

COPY public.education (summary_id, institution, speciality, graduated, type) FROM stdin;
\.


--
-- Data for Name: experience; Type: TABLE DATA; Schema: public; Owner: huvalk
--

COPY public.experience (company_name, role, responsibilities, start, stop, summary_id) FROM stdin;
\.


--
-- Data for Name: favorite; Type: TABLE DATA; Schema: public; Owner: huvalk
--

COPY public.favorite (user_id, favorite_id) FROM stdin;
\.


--
-- Data for Name: organization; Type: TABLE DATA; Schema: public; Owner: huvalk
--

COPY public.organization (id, name, site, about) FROM stdin;
\.


--
-- Data for Name: person; Type: TABLE DATA; Schema: public; Owner: huvalk
--

COPY public.person (id, name, gender, birthday) FROM stdin;
\.


--
-- Data for Name: requirement; Type: TABLE DATA; Schema: public; Owner: huvalk
--

COPY public.requirement (summary_id, driver_license, has_car, schedule, employment) FROM stdin;
\.


--
-- Data for Name: response; Type: TABLE DATA; Schema: public; Owner: huvalk
--

COPY public.response (summary_id, vacancy_id, rejected, approved, date) FROM stdin;
\.


--
-- Data for Name: session; Type: TABLE DATA; Schema: public; Owner: huvalk
--

COPY public.session (user_id, session_id, expires) FROM stdin;
\.


--
-- Data for Name: summary; Type: TABLE DATA; Schema: public; Owner: huvalk
--

COPY public.summary (id, author, keywords) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: huvalk
--

COPY public.users (id, login, password, organization_id, person_id, tag, email, phone, registered, avatar) FROM stdin;
\.


--
-- Data for Name: vacancy; Type: TABLE DATA; Schema: public; Owner: huvalk
--

COPY public.vacancy (id, organization_id, name, description, with_tax, responsibilities, conditions, keywords, salary_from, salary_to) FROM stdin;
\.


--
-- Name: organization_id_seq; Type: SEQUENCE SET; Schema: public; Owner: huvalk
--

SELECT pg_catalog.setval('public.organization_id_seq', 1, false);


--
-- Name: person_id_seq; Type: SEQUENCE SET; Schema: public; Owner: huvalk
--

SELECT pg_catalog.setval('public.person_id_seq', 1, false);


--
-- Name: summary_id_seq; Type: SEQUENCE SET; Schema: public; Owner: huvalk
--

SELECT pg_catalog.setval('public.summary_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: huvalk
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: vacancy_id_seq; Type: SEQUENCE SET; Schema: public; Owner: huvalk
--

SELECT pg_catalog.setval('public.vacancy_id_seq', 1, false);


--
-- Name: organization organization_pkey; Type: CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.organization
    ADD CONSTRAINT organization_pkey PRIMARY KEY (id);


--
-- Name: person person_pkey; Type: CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.person
    ADD CONSTRAINT person_pkey PRIMARY KEY (id);


--
-- Name: summary summary_pkey; Type: CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.summary
    ADD CONSTRAINT summary_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: vacancy vacancy_pkey; Type: CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.vacancy
    ADD CONSTRAINT vacancy_pkey PRIMARY KEY (id);


--
-- Name: education education_summary_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.education
    ADD CONSTRAINT education_summary_id_fkey FOREIGN KEY (summary_id) REFERENCES public.summary(id) ON DELETE CASCADE;


--
-- Name: experience experience_summary_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.experience
    ADD CONSTRAINT experience_summary_id_fkey FOREIGN KEY (summary_id) REFERENCES public.summary(id) ON DELETE CASCADE;


--
-- Name: favorite favorite_favorite_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.favorite
    ADD CONSTRAINT favorite_favorite_id_fkey FOREIGN KEY (favorite_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: favorite favorite_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.favorite
    ADD CONSTRAINT favorite_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: requirement requirement_summary_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.requirement
    ADD CONSTRAINT requirement_summary_id_fkey FOREIGN KEY (summary_id) REFERENCES public.summary(id);


--
-- Name: response response_summary_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.response
    ADD CONSTRAINT response_summary_id_fkey FOREIGN KEY (summary_id) REFERENCES public.summary(id);


--
-- Name: response response_vacancy_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.response
    ADD CONSTRAINT response_vacancy_id_fkey FOREIGN KEY (vacancy_id) REFERENCES public.vacancy(id);


--
-- Name: session session_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.session
    ADD CONSTRAINT session_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: summary summary_author_fkey; Type: FK CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.summary
    ADD CONSTRAINT summary_author_fkey FOREIGN KEY (author) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: users users_organization_fkey; Type: FK CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_organization_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id);


--
-- Name: users users_person_fkey; Type: FK CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_person_fkey FOREIGN KEY (person_id) REFERENCES public.person(id);


--
-- Name: vacancy vacancy_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: huvalk
--

ALTER TABLE ONLY public.vacancy
    ADD CONSTRAINT vacancy_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

