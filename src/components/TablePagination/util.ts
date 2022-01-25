// Logic for showing visible page buttons based on input of current page and total page
// Based on USWDS guidelines - https://designsystem.digital.gov/components/pagination/

// don't change this value; logic depends on it
type MAX_SLOTS = 7;
// eslint-disable-next-line @typescript-eslint/no-redeclare, no-redeclare
export const MAX_SLOTS: MAX_SLOTS = 7;

// TODO - see if I can get tuple working with MAX_SLOTS type
// see https://gist.github.com/DylanSp/b6eea0dbf9a6d176a9c8125cc8fa347b

// Creates an array of page numbers starting a 1, from an input of a number
const fillMinPages = (totalPages: number) => {
  return Array.from({ length: totalPages }, (_, i) => i + 1);
};

type PageNumber = {
  is: 'pageNumber';
  pageNumber: number;
  isCurrent: boolean;
};

type OverflowIndicator = {
  is: 'overflowIndicator';
  slotNumber: 2 | 6;
};

type PaginationSlot = PageNumber | OverflowIndicator;

const fillMinSlots = (totalPages: number): PageNumber[] => {
  return Array.from({ length: totalPages }, (_, i) => ({
    is: 'pageNumber',
    pageNumber: i + 1,
    isCurrent: false
  }));
};

export const getPaginationSlots = (
  currentPage: number,
  totalPages: number
): PaginationSlot[] => {
  if (totalPages <= MAX_SLOTS) {
    const slots = fillMinSlots(totalPages);
    slots[currentPage - 1].isCurrent = true;
    return slots;
  }

  const slots: PaginationSlot[] = [];
  slots[0] = {
    is: 'pageNumber',
    pageNumber: 1,
    isCurrent: currentPage === 1
  };

  if (currentPage > 4) {
    slots[1] = {
      is: 'overflowIndicator',
      slotNumber: 2
    };
    slots[2] = {
      is: 'pageNumber',
      pageNumber: currentPage - 1,
      isCurrent: false
    };
  } else {
    slots[1] = {
      is: 'pageNumber',
      pageNumber: 2,
      isCurrent: currentPage === 2
    };
    slots[2] = {
      is: 'pageNumber',
      pageNumber: 3,
      isCurrent: currentPage === 3
    };
  }

  /*
  slots[2] = {
    is: 'pageNumber',
    pageNumber: currentPage - 1,
    isCurrent: false
  };
  */

  slots[3] = {
    is: 'pageNumber',
    pageNumber: currentPage,
    isCurrent: true
  };

  /*
  slots[4] = {
    is: 'pageNumber',
    pageNumber: currentPage + 1,
    isCurrent: false
  };
  */

  if (totalPages - currentPage > 3) {
    slots[MAX_SLOTS - 3] = {
      is: 'pageNumber',
      pageNumber: currentPage + 1,
      isCurrent: false
    };

    slots[MAX_SLOTS - 2] = {
      is: 'overflowIndicator',
      slotNumber: 6
    };
  } else {
    slots[MAX_SLOTS - 3] = {
      is: 'pageNumber',
      pageNumber: MAX_SLOTS - 2,
      isCurrent: currentPage === MAX_SLOTS - 2
    };
    slots[MAX_SLOTS - 2] = {
      is: 'pageNumber',
      pageNumber: MAX_SLOTS - 1,
      isCurrent: currentPage === MAX_SLOTS - 1
    };
  }

  slots[MAX_SLOTS - 1] = {
    is: 'pageNumber',
    pageNumber: totalPages,
    isCurrent: currentPage === totalPages
  };

  return slots;
};

const getVisiblePages = (currentPage: number, totalPages: number) => {
  // If total pages are less than 6 (minimum for elipsis), create an array of available page based on total
  if (totalPages <= MAX_SLOTS) {
    return fillMinPages(totalPages);
  }
  // If total is more than the minimum for elipsis (6) and IS NOT adjacent to the start or end of pages
  // Returns an array with the start page, current pages, two adjacent pages, and end page
  if (currentPage % 4 >= 0 && currentPage > 3 && currentPage + 2 < totalPages) {
    return [1, currentPage - 1, currentPage, currentPage + 1, totalPages];
  }
  // If total is more than the minimum for elipsis (6) and IS adjacent to the end page within 3 pages
  // Returns an array with the start page, and the last four pages
  if (
    currentPage % 4 >= 0 &&
    currentPage > 3 &&
    currentPage + 2 >= totalPages
  ) {
    return [1, totalPages - 3, totalPages - 2, totalPages - 1, totalPages];
  }
  // If total is more than the minimum for elipsis (6) and IS adjacent to the start page within 3 pages
  // Returns an array with the first four pages and end page
  return [1, 2, 3, 4, totalPages];
};

export default getVisiblePages;
