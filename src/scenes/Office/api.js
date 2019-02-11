import { getClient, checkResponse } from 'shared/Swagger/api';

// MOVE QUEUE
export async function RetrieveMovesForOffice(queueType) {
  const client = await getClient();
  const response = await client.apis.queues.showQueue({
    queueType,
  });
  checkResponse(response, 'failed to retrieve moves due to server error');
  return response.body;
}

// MOVE
export async function LoadMove(moveId) {
  const client = await getClient();
  const response = await client.apis.moves.showMove({
    moveId,
  });
  checkResponse(response, 'failed to load move due to server error');
  return response.body;
}

// HHG invoice
export async function SendHHGInvoice(shipmentId) {
  const client = await getClient();
  const response = await client.apis.shipments.createAndSendHHGInvoice({
    shipmentId,
  });
  checkResponse(response, 'failed to send invoice to server error');
  return response.body;
}

// PPM attachments
export async function DownloadPPMAttachments(ppmId, docTypes) {
  const client = await getClient();
  const response = await client.apis.ppm.createPPMAttachments({
    personallyProcuredMoveId: ppmId,
    docTypes: docTypes,
  });
  checkResponse(response, 'failed to create PPM attachments PDF');
  return response.body;
}
