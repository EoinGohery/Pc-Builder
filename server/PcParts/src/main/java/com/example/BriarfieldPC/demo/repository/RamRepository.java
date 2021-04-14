package com.example.BriarfieldPC.demo.repository;

import com.example.BriarfieldPC.demo.entity.Ram;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface RamRepository extends JpaRepository<Ram,Long> {
    List<Ram> findAll();
}
