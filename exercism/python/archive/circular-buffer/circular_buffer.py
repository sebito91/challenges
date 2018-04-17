""" Mdolue to implement a circular-buffer """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

class BufferFullException(Exception):
    """ define execption when buffer is full """
    def __init__(self, message=None):
        if not message:
            message = "buffer is full"
        super(BufferFullException, self).__init__(message)

class BufferEmptyException(Exception):
    """ define exception when buffer is empty """
    def __init__(self, message=None):
        if not message:
            message = "buffer is empty"
        super(BufferEmptyException, self).__init__(message)

class CircularBuffer(object):
    """ definition of the back CircularBuffer class """
    def __init__(self, datasize):
        if datasize <= 0:
            raise ValueError("improper size for CircularBuffer: {}".format(datasize))

        self.buffer = [None] * datasize
        self.capacity = datasize
        self.current = (0, 0)

    def get_elem(self, index):
        """ helper function to increment counters """
        temp = self.current[0]

        if index == 0:
            self.current = ((self.current[0] + 1) % (self.capacity), self.current[1])
        else:
            temp = self.current[1]
            self.current = (self.current[0], (self.current[1] + 1) % (self.capacity))

        return temp

    def read(self):
        """ read function as part of CircularBuffer """
        if len(self.buffer) < 1 or all(each is None for each in self.buffer):
            raise BufferEmptyException("tried reading from empty buffer")

        idx = self.get_elem(0)
        data = self.buffer[idx]
        self.buffer[idx] = None

        return data

    def write(self, data):
        """ write function as part of CircularBuffer """
        if self.current[0] == self.current[1] and self.buffer[self.current[0]]:
            raise BufferFullException("cannot add {} to full buffer".format(data))
        self.buffer[self.get_elem(1)] = data

    def overwrite(self, data):
        """ overwrite the oldest data first """
        self.buffer[self.get_elem(0)] = data

    def clear(self):
        """ clear out the buffer """
        self.buffer = [None] * self.capacity
