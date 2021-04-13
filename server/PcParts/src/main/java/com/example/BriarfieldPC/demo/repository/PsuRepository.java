package com.example.BriarfieldPC.demo.repository;

import com.example.BriarfieldPC.demo.entity.PSU;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface PsuRepository extends JpaRepository<PSU,Long> {
    List<PSU> findAll();
}
