-- +goose Up
-- +goose StatementBegin

CREATE TABLE sessions (
                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          clinic_id UUID NOT NULL REFERENCES clinics(id),
                          patient_id UUID NOT NULL REFERENCES patients(id),
                          dt_session TIMESTAMP NOT NULL,
                          paid BOOLEAN DEFAULT false,
                          confirmed BOOLEAN DEFAULT false,
                          value NUMERIC(10, 2) NOT NULL,
                          report TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd
