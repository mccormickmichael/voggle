import os

class Board(object):
    def __init__(self, rows, cols, cells):
        self.rows = rows
        self.cols = cols
        self.cells = cells
        self._build_matrix()

    def _build_matrix(self):
        matrix = []
        for row_index in range(self.rows):
            row = []
            for col_index in range(self.cols):
                cell_index = (row_index * self.cols + col_index) % len(self.cells)
                row.append(Cell(row_index,
                                col_index,
                                str(self.cells[cell_index])))
            matrix.append(row)
        self.matrix = matrix
            

    def __str__(self):
        col_sizes = [self._max_col_len(c) for c in range(self.cols)]
        
        row_separator = ('-' * (sum(col_sizes) + 3 * self.cols + 1)) + os.linesep
        result = row_separator
        for row in self.matrix:
            for cell_index in range(self.cols):
                cell = row[cell_index]
                pad = col_sizes[cell_index] - len(cell)
                result += '| ' + (' ' * pad) + str(cell) + ' '
            result += '|' + os.linesep
            result += row_separator
        return result
        
    def _max_col_len(self, col_index):
        return max( ( len(r[col_index]) for r in self.matrix) )


class Cell(object):
    def __init__(self, row_index, col_index, value):
        self.row_index = row_index
        self.col_index = col_index
        self.value = value

    def __str__(self):
        return self.value

    def __len__(self):
        return len(self.value)
