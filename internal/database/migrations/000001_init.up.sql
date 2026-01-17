-- Группа мышц
CREATE TABLE muscle_groups (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT      NOT NULL,
)

-- Тренажеры
CREATE TABLE machines (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT      NOT NULL,
    description TEXT,
    muscle_group_id BIGINT REFERENCES muscle_group(id) ON DELETE SET NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
)

CREATE TABLE machines_images (
    id   BIGSERIAL PRIMARY KEY,
    machine_id BIGINT REFERENCES machines(id) ON DELETE SET CASCADE,
    image_url VARCHAR(500) NOT NULL,
    alt_text VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
)

-- Упражнения
CREATE TABLE exercises (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT      NOT NULL,
    description TEXT,
    warning TEXT,
    recommended_reps_amount SMALLINT,
    recommended_rest_duration INT,
    is_gym_required BOOLEAN,
    tutorial_video_url VARCHAR(500),
    machine_id BIGINT REFERENCES machines(id) ON DELETE SET NULL,
    muscle_group_id BIGINT REFERENCES muscle_group(id) ON DELETE SET NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
);

CREATE TABLE exercises_images (
    id   BIGSERIAL PRIMARY KEY,
    exercise_id BIGINT REFERENCES exercises(id) ON DELETE SET CASCADE,
    image_url VARCHAR(500) NOT NULL,
    alt_text VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
)