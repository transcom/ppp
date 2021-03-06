-- The US Navy is not going to connect to the Orders Gateway directly in the
-- near-term. Instead, they are providing us flat-file exports from their
-- orders database, and we run those flat-file exports through the Navy Orders
-- Muncher (nom) which handles converting the flat file into the appropriate
-- JSON structures and uploading them to the Orders API.
-- The Orders API uses client certificate authentication. Only certificates
-- signed by a trusted CA (such as DISA) are allowed. As nom is a command-line
-- tool run by a person, using that person's CAC as the certificate is a
-- convenient way to permit a single trusted individual to upload Orders.
-- Once the Navy completes their integration between NSIPS and the Orders API,
-- this CAC certificate should be removed.
-- Alternatively, if nom is integrated into MilMove and nom or the Orders API
-- gains the ability to authenticate using login.gov, then we should use that
-- instead for this particular Navy use-case and remove this CAC certificate.
-- In development, load no new certs. This is just for the real environments.
