import { get } from 'lodash';
import { swaggerRequest } from 'shared/Swagger/request';
import { getClient } from 'shared/Swagger/api';

const approvePpmLabel = 'PPMs.approvePPM';

export function approvePPM(personallyProcuredMoveId) {
  const label = approvePpmLabel;
  const swaggerTag = 'office.approvePPM';
  return swaggerRequest(getClient, swaggerTag, { personallyProcuredMoveId }, { label });
}

export function loadPPMs(moveId) {
  const label = 'office.loadPPMs';
  const swaggerTag = 'ppm.indexPersonallyProcuredMoves';
  return swaggerRequest(getClient, swaggerTag, { moveId }, { label });
}

export function selectPPM(state) {
  // Note: will need to be changed when we support multiple PPMS
  const ppmId = Object.keys(get(state, 'entities.personallyProcuredMove', {}))[0];
  return get(state, `entities.personallyProcuredMove.${ppmId}`, {});
}

export function updatePPM(
  moveId,
  personallyProcuredMoveId,
  payload /*shape: {size, weightEstimate, estimatedIncentive}*/,
) {
  const label = 'office.updatePPM';
  const swaggerTag = 'ppm.patchPersonallyProcuredMove';
  return swaggerRequest(
    getClient,
    swaggerTag,
    {
      moveId,
      personallyProcuredMoveId: personallyProcuredMoveId,
      patchPersonallyProcuredMovePayload: payload,
    },
    { label },
  );
}
