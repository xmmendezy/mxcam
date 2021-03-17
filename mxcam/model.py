# -*- coding: utf-8 -*-

from datetime import datetime
from sqlalchemy import create_engine
from sqlalchemy.orm import scoped_session, sessionmaker, relationship
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, String, DateTime, ForeignKey

__all__ = [
    'init_db', 'User', 'Source', 'Ip', 'Command', 'Log',
]

engine = create_engine('sqlite:///database.db')
session = scoped_session(sessionmaker(engine))
BaseModel = declarative_base()
BaseModel.query = session.query_property()


def init_db():
    BaseModel.metadata.create_all(engine)


class Model(BaseModel):
    __abstract__ = True

    created_on = Column(DateTime, default=datetime.utcnow)
    updated_on = Column(DateTime, default=datetime.utcnow, onupdate=datetime.utcnow)

    def save(self):
        session.add(self)
        session.commit()

    def delete(self):
        session.delete(self)
        session.commit()


class User(Model):
    __tablename__ = 'user'

    id = Column(Integer, primary_key=True)
    username = Column(String(200))
    password = Column(String(200))

    def __init__(self, username, password):
        self.username = username
        self.password = password


class Source(Model):
    __tablename__ = 'source'

    id = Column(Integer, primary_key=True)
    name = Column(String(200))
    comment = Column(String(200))
    ips = relationship("Ip", cascade="all, delete-orphan")

    def __init__(self, name, comment, ips=[]):
        self.name = name
        self.comment = comment
        self.ips = ips


class Ip(Model):
    __tablename__ = 'ip'

    id = Column(Integer, primary_key=True)
    ip = Column(String(200))
    port = Column(String(200))
    source = relationship('Source', uselist=False)
    source_id = Column(Integer, ForeignKey('source.id'), nullable=False)

    def __init__(self, ip, port, source_id):
        self.ip = ip
        self.port = port
        self.source_id = source_id

class Command(Model):
    __tablename__ = 'command'

    id = Column(Integer, primary_key=True)
    name = Column(String(200))
    path = Column(String(200))
    value = Column(String(10000))
    message_logs = Column(String(10000))
    default_message = Column(String(400))

    def __init__(self, name, path, value, message_logs = '{}', default_message = ''):
        self.name = name
        self.path = path
        self.value = value
        self.message_logs = message_logs
        self.default_message = default_message

class Log(Model):
    __tablename__ = 'log'

    id = Column(Integer, primary_key=True)
    value = Column(String(100000))

    def __init__(self, value):
        self.value = value
