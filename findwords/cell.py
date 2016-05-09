

class Cell(object):

    def __init__(self, row, col, value):
        self.value = value
        self.row = row
        self.col = col
        self._neighbors = []

    def neighbors(self):
        return self._neighbors

    def _resolve_neighbors(self, board):
        self._neighbors = []
        rel = [(self.row+1, self.col-1), (self.row+1, self.col+0), (self.row+1, self.col+1),
               (self.row+0, self.col-1),                           (self.row+0, self.col+1),
               (self.row-1, self.col-1), (self.row-1, self.col+0), (self.row-1, self.col+1)]
        for row, col in rel:
            if row in range(board.board_size()) and col in range(board.board_size()):
                self._neighbors.append(board[row][col])

    def __repr__(self):
        return '[{}, {}]:{}'.format(self.row, self.col, self.value)


class Board(object):
    def __init__(self, values):
        self._values = values
        self._board_size = 0
        self._board = []

        self._discover_board_size()
        self._build_board()

    def board_size(self):
        return self._board_size

    def cells(self):
        return [col for row in self._board for col in row]

    def board(self):
        return self._board

    def __getitem__(self, index):
        return self._board[index]

    def _discover_board_size(self):
        sizes = {x*x: x for x in range(1, 6)}
        self._board_size = sizes.get(len(self._values), 0)
        if self._board_size is 0:
            raise ValueError('Board must be a square between 2x2 and 6x6')

    def _build_board(self):
        self._build_cells()
        self._resolve_neighbors()

    def _build_cells(self):
        # TODO: there must be a more pythonic way to write this...
        self._board = []
        stream = (v for v in self._values)
        for row_num in range(self._board_size):
            row = []
            self._board.append(row)
            for col_num in range(self._board_size):
                row.append(Cell(row_num, col_num, next(stream)))

    def _resolve_neighbors(self):
        for cell in self.cells():
            cell._resolve_neighbors(self)
