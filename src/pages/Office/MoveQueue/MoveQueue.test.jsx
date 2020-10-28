import React from 'react';
import { mount } from 'enzyme';

import MoveQueue from './MoveQueue';

import { MockProviders } from 'testUtils';

jest.mock('hooks/queries', () => ({
  useMovesQueueQueries: () => {
    return {
      isLoading: false,
      isError: false,
      queueMovesResult: {
        totalCount: 2,
        queueMoves: [
          {
            id: 'move1',
            customer: {
              first_name: 'test first',
              last_name: 'test last',
              dodID: '555555555',
            },
            locator: 'AB5P',
            departmentIndicator: 'ARMY',
            shipmentsCount: 2,
            status: 'NEW',
            destinationDutyStation: {
              name: 'Area 51',
            },
            originGBLOC: 'EEEE',
          },
          {
            id: 'move2',
            customer: {
              first_name: 'test another first',
              last_name: 'test another last',
              dodID: '4444444444',
            },
            locator: 'T12A',
            departmentIndicator: 'NAVY_AND_MARINES',
            shipmentsCount: 1,
            status: 'APPROVED',
            destinationDutyStation: {
              name: 'Los Alamos',
            },
            originGBLOC: 'EEEE',
          },
        ],
      },
    };
  },
}));

describe('MoveQueue', () => {
  const wrapper = mount(
    <MockProviders initialEntries={['moves/queue']}>
      <MoveQueue />
    </MockProviders>,
  );

  it('should render the h1', () => {
    expect(wrapper.find('h1').text()).toBe('All moves (2)');
  });

  it('should render the table', () => {
    expect(wrapper.find('Table').exists()).toBe(true);
  });

  it('should format the column data', () => {
    const moves = wrapper.find('tbody tr');

    const firstMove = moves.at(0);
    expect(firstMove.find({ 'data-testid': 'name-0' }).text()).toBe('test last, test first');
    expect(firstMove.find({ 'data-testid': 'customer.dodID-0' }).text()).toBe('555555555');
    expect(firstMove.find({ 'data-testid': 'status-0' }).text()).toBe('NEW');
    expect(firstMove.find({ 'data-testid': 'locator-0' }).text()).toBe('AB5P');
    expect(firstMove.find({ 'data-testid': 'branch-0' }).text()).toBe('Army');
    expect(firstMove.find({ 'data-testid': 'shipmentsCount-0' }).text()).toBe('2');
    expect(firstMove.find({ 'data-testid': 'destinationDutyStation.name-0' }).text()).toBe('Area 51');
    expect(firstMove.find({ 'data-testid': 'originGBLOC-0' }).text()).toBe('EEEE');

    const secondMove = moves.at(1);
    expect(secondMove.find({ 'data-testid': 'name-1' }).text()).toBe('test another last, test another first');
    expect(secondMove.find({ 'data-testid': 'customer.dodID-1' }).text()).toBe('4444444444');
    expect(secondMove.find({ 'data-testid': 'status-1' }).text()).toBe('APPROVED');
    expect(secondMove.find({ 'data-testid': 'locator-1' }).text()).toBe('T12A');
    expect(secondMove.find({ 'data-testid': 'branch-1' }).text()).toBe('Navy and Marine Corps');
    expect(secondMove.find({ 'data-testid': 'shipmentsCount-1' }).text()).toBe('1');
    expect(secondMove.find({ 'data-testid': 'destinationDutyStation.name-1' }).text()).toBe('Los Alamos');
    expect(secondMove.find({ 'data-testid': 'originGBLOC-1' }).text()).toBe('EEEE');
  });
});
