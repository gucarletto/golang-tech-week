CREATE TABLE videos (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    file_path VARCHAR(255) NOT NULL,
    hls_path VARCHAR(255),
    manifest_path VARCHAR(255),
    s3_manifest_path VARCHAR(255),
    s3_url VARCHAR(255),
    status VARCHAR(50) NOT NULL,
    upload_status VARCHAR(50),
    error_message TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
