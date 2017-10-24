package com.neunn.cloud.suyan.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.neunn.cloud.suyan.entity.User;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
}
