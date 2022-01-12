import * as fc from 'fast-check';

import getVisiblePages, { getPaginationSlots, MAX_SLOTS } from './util';

describe('TablePagination Util', () => {
  it('returns an empty array', () => {
    const currentPage: number = 0;
    const totalPages: number = 0;
    const expectedPages: number[] = [];
    expect(getVisiblePages(currentPage, totalPages)).toEqual(expectedPages);
  });

  it('returns an array of only 3 pages', () => {
    const currentPage: number = 1;
    const totalPages: number = 3;
    const expectedPages: number[] = [1, 2, 3];
    expect(getVisiblePages(currentPage, totalPages)).toEqual(expectedPages);
  });

  it('returns an array of only first 4 pages and max page', () => {
    const currentPage: number = 0;
    const totalPages: number = 60;
    const expectedPages: number[] = [1, 2, 3, 4, 60];
    expect(getVisiblePages(currentPage, totalPages)).toEqual(expectedPages);
  });

  it('returns an array with starting page, max page, current page, and two surrounding current', () => {
    const currentPage: number = 25;
    const totalPages: number = 80;
    const expectedPages: number[] = [1, 24, 25, 26, 80];
    expect(getVisiblePages(currentPage, totalPages)).toEqual(expectedPages);
  });

  it('returns an array with starting page, current page-1, and current page through max page', () => {
    const currentPage: number = 48;
    const totalPages: number = 50;
    const expectedPages: number[] = [1, 47, 48, 49, 50];
    expect(getVisiblePages(currentPage, totalPages)).toEqual(expectedPages);
  });

  it('returns an array with starting page and current page with 3 preceding pages', () => {
    const currentPage: number = 50;
    const totalPages: number = 50;
    const expectedPages: number[] = [1, 47, 48, 49, 50];
    expect(getVisiblePages(currentPage, totalPages)).toEqual(expectedPages);
  });
});

describe.only('getPaginationSlots property-based testing', () => {
  it('Removes extra slots if there are fewer than MAX_SLOTS', () => {
    fc.assert(
      fc.property(
        fc.integer({ min: 1, max: MAX_SLOTS - 1 }),
        fc.integer({ min: 1, max: MAX_SLOTS - 1 }),
        (currentPage, totalPages) => {
          fc.pre(currentPage <= totalPages);
          const slots = getPaginationSlots(currentPage, totalPages);
          expect(slots).toHaveLength(totalPages);
        }
      )
    );
  });

  it('Always highlights the current page', () => {
    fc.assert(
      fc.property(
        fc.integer({ min: 1 }),
        fc.integer({ min: 1 }),
        // fc.context(),
        (currentPage, totalPages /* , ctx */) => {
          fc.pre(currentPage <= totalPages);
          // ctx.log(`currentPage: ${currentPage}, totalPages: ${totalPages}`);
          const slots = getPaginationSlots(currentPage, totalPages);
          const currentPageSlot = slots.find(
            slot =>
              slot.is === 'pageNumber' &&
              slot.pageNumber === currentPage &&
              slot.isCurrent
          );
          expect(currentPageSlot).not.toBeUndefined();
        }
      )
    );
  });

  it('Never shows more slots than MAX_SLOTS', () => {
    fc.assert(
      fc.property(
        fc.integer({ min: 1 }),
        fc.integer({ min: 1 }),
        (currentPage, totalPages) => {
          fc.pre(currentPage <= totalPages);
          const slots = getPaginationSlots(currentPage, totalPages);
          expect(slots.length).toBeLessThanOrEqual(MAX_SLOTS);
        }
      )
    );
  });

  it('Always shows the previous page adjacent to the current page', () => {
    fc.assert(
      // start at 2 so there exists a previous page
      fc.property(
        fc.integer({ min: 2 }),
        fc.integer({ min: 2 }),
        (currentPage, totalPages) => {
          fc.pre(currentPage <= totalPages);
          const slots = getPaginationSlots(currentPage, totalPages);
          const previousPage = slots.find(
            slot =>
              slot.is === 'pageNumber' && slot.pageNumber === currentPage - 1
          );
          expect(previousPage).not.toBeUndefined();
        }
      )
    );
  });

  it('Always shows the next page adjacent to the current page', () => {
    fc.assert(
      fc.property(
        fc.integer({ min: 1 }),
        fc.integer({ min: 1 }),
        (currentPage, totalPages) => {
          fc.pre(currentPage < totalPages);
          const slots = getPaginationSlots(currentPage, totalPages);
          const previousPage = slots.find(
            slot =>
              slot.is === 'pageNumber' && slot.pageNumber === currentPage + 1
          );
          expect(previousPage).not.toBeUndefined();
        }
      )
    );
  });

  it('Shows an overflow indicator in second slot when there are pages missing between current page (and neighbors) and beginning', () => {
    fc.assert(
      // start currentPage at 4 so there will be pages missing
      // start totalPages at MAX_SLOTS + 1 so there will be overflow
      fc.property(
        fc.integer({ min: 5 }),
        fc.integer({ min: MAX_SLOTS + 1 }),
        // fc.context(),
        (currentPage, totalPages /* , ctx */) => {
          fc.pre(currentPage <= totalPages);

          const slots = getPaginationSlots(currentPage, totalPages);
          // ctx.log(slots);
          expect(slots[1].is).toEqual('overflowIndicator');
        }
      )
    );
  });

  it("Doesn't show overflow indicator in second slot if indicator could be replaced with page number and remain continuous", () => {
    fc.assert(
      fc.property(
        fc.integer({ min: MAX_SLOTS + 1 }),
        // fc.context(),
        (totalPages /* , ctx */) => {
          const currentPage = 4;
          const slots = getPaginationSlots(currentPage, totalPages);
          // ctx.log(slots);
          expect(slots[1].is).toEqual('pageNumber');
        }
      )
    );
  });

  it('Shows an overflow indicator in second-to-last slot when there are pages missing between current page (and neighbors) and end', () => {
    fc.assert(
      // start totalPages at MAX_SLOTS + 1 so there will be overflow
      fc.property(
        fc.integer({ min: 1 }),
        fc.integer({ min: MAX_SLOTS + 1 }),
        (currentPage, totalPages) => {
          fc.pre(totalPages - currentPage > 4);

          const slots = getPaginationSlots(currentPage, totalPages);
          expect(slots[slots.length - 2].is).toEqual('overflowIndicator');
        }
      )
    );
  });

  it("Doesn't show overflow indicator in second-to-last slot if indicator could be replaced with page number and remain continuous", () => {
    fc.assert(
      fc.property(fc.integer({ min: MAX_SLOTS + 1 }), totalPages => {
        const currentPage = totalPages - 3;
        const slots = getPaginationSlots(currentPage, totalPages);
        expect(slots[slots.length - 2].is).toEqual('pageNumber');
      })
    );
  });

  it('Never shows overflow indicators when totalPages <= MAX_SLOTS', () => {
    fc.assert(
      fc.property(
        fc.integer({ min: 1, max: MAX_SLOTS }),
        fc.integer({ min: 1, max: MAX_SLOTS }),
        (currentPage, totalPages) => {
          fc.pre(currentPage <= totalPages);

          const slots = getPaginationSlots(currentPage, totalPages);
          expect(slots.every(slot => slot.is === 'pageNumber')).toBeTruthy();
        }
      )
    );
  });

  it('Never shows overflow indicators outside of second and second-to-last slots when totalPages > MAX_SLOTS', () => {
    fc.assert(
      fc.property(
        fc.integer({ min: 1 }),
        fc.integer({ min: MAX_SLOTS + 1 }),
        (currentPage, totalPages) => {
          fc.pre(currentPage <= totalPages);

          const slots = getPaginationSlots(currentPage, totalPages);
          expect(
            slots.every(
              (slot, index) =>
                index === 1 ||
                index === slots.length - 2 ||
                slot.is === 'pageNumber'
            )
          ).toBeTruthy();

          /*
        for (let i = 0; i < slots.length; i++) {
          if (i === 1 || i === slots.length - 2) {
            continue;
          }
          expect(slots[i].is).not.toEqual('overflowIndicator');
        }
        */
        }
      )
    );
  });
});

describe.skip('getPaginationSlots', () => {
  it('Returns an empty array for 0 total pages', () => {
    const totalPages = 0;

    const slots = getPaginationSlots(0, totalPages);

    expect(slots).toHaveLength(0);
  });

  it('Returns an array with length = totalPages if totalPages < MAX_SLOTS', () => {
    const totalPages = MAX_SLOTS - 1;

    const slots = getPaginationSlots(0, totalPages);

    expect(slots).toHaveLength(totalPages);
  });

  it('Shows an overflow indicator in second-to-last slot when currentPage is first page and totalPages > MAX_SLOTS', () => {
    const currentPage = 1;
    const totalPages = MAX_SLOTS + 1;

    const slots = getPaginationSlots(currentPage, totalPages);

    expect(slots).toHaveLength(MAX_SLOTS);
    expect(slots[slots.length - 2]).toEqual({ is: 'overflowIndicator' });
  });

  it('Shows overflow indicators in second and second-to-last slots when currentPage is in middle of range and totalPages > MAX_SLOTS', () => {
    const currentPage = 50;
    const totalPages = 100;

    const slots = getPaginationSlots(currentPage, totalPages);

    expect(slots).toHaveLength(MAX_SLOTS);
    expect(slots[1]).toEqual({ is: 'overflowIndicator' });
    expect(slots[slots.length - 2]).toEqual({ is: 'overflowIndicator' });
  });

  it('Shows an overflow indicator in second slot when currentPage is last page and totalPages > MAX_SLOTS', () => {
    const currentPage = MAX_SLOTS + 1;
    const totalPages = MAX_SLOTS + 1;

    const slots = getPaginationSlots(currentPage, totalPages);

    expect(slots).toHaveLength(MAX_SLOTS);
    expect(slots[1]).toEqual({ is: 'overflowIndicator' });
  });
});
